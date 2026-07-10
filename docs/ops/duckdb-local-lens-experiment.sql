-- DuckDB local read-only lens experiment for Issue #312.
-- The runner sets trace_path, motion_path, and catalog_path before loading this file.
-- Every statement uses temporary in-memory views; the final result contains aggregates only.

SET autoinstall_known_extensions = false;
SET autoload_known_extensions = false;

CREATE TEMP VIEW trace_rows AS
SELECT *
FROM read_csv_auto(getvariable('trace_path'), header = true);

CREATE TEMP VIEW motion_rows AS
SELECT *
FROM read_csv_auto(getvariable('motion_path'), header = true);

CREATE TEMP VIEW catalog_rows AS
SELECT *
FROM read_json_auto(getvariable('catalog_path'));

CREATE TEMP VIEW trace_dates AS
SELECT try_cast("日期" AS DATE) AS trace_date
FROM trace_rows;

CREATE TEMP VIEW trace_duplicate_counts AS
SELECT
  count(*) AS copies,
  coalesce(trim("色散内容"), '') = '' AS has_blank_content
FROM trace_rows
GROUP BY
  "日期",
  "色散内容",
  "关联项目",
  "判断线索",
  "完成质量",
  "情绪",
  "执行情境强度",
  "状态",
  "色散类型",
  "认知挑战";

CREATE TEMP VIEW catalog_parts AS
SELECT
  part.path AS path,
  part.start AS start_id,
  part."end" AS end_id,
  part.count AS item_count
FROM catalog_rows,
UNNEST(parts) AS unpacked(part);

CREATE TEMP VIEW combined_csv_rows AS
SELECT
  CASE
    WHEN filename LIKE '%色散 样本.csv' THEN 'trace'
    WHEN filename LIKE '%探筹 样本.csv' THEN 'motion'
    ELSE 'unexpected'
  END AS source_kind,
  * EXCLUDE (filename)
FROM read_csv_auto(
  [getvariable('trace_path'), getvariable('motion_path')],
  header = true,
  union_by_name = true,
  filename = true
);

-- Q1: How many logical rows are present in each source, and did Chinese content survive?
SELECT
  'Q1_source_overview' AS question_id,
  'Logical row counts and Chinese-content checks' AS finding,
  json_object(
    'trace_rows', (SELECT count(*) FROM trace_rows),
    'motion_rows', (SELECT count(*) FROM motion_rows),
    'catalog_rows', (SELECT count(*) FROM catalog_rows),
    'trace_rows_with_chinese', (
      SELECT count(*)
      FROM trace_rows
      WHERE regexp_matches(coalesce("色散内容", ''), '[一-龥]')
    ),
    'motion_rows_with_chinese', (
      SELECT count(*)
      FROM motion_rows
      WHERE regexp_matches(coalesce("动作", ''), '[一-龥]')
    )
  ) AS details

UNION ALL

-- Q2: Which fields are blank, and are there duplicate rows or duplicate candidate keys?
SELECT
  'Q2_data_quality' AS question_id,
  'Blank-field and duplicate diagnostics without returning private text' AS finding,
  json_object(
    'trace_blank_counts', (
      SELECT json_object(
        '日期', count(*) FILTER (WHERE coalesce(trim(cast("日期" AS VARCHAR)), '') = ''),
        '色散内容', count(*) FILTER (WHERE coalesce(trim(cast("色散内容" AS VARCHAR)), '') = ''),
        '关联项目', count(*) FILTER (WHERE coalesce(trim(cast("关联项目" AS VARCHAR)), '') = ''),
        '判断线索', count(*) FILTER (WHERE coalesce(trim(cast("判断线索" AS VARCHAR)), '') = ''),
        '完成质量', count(*) FILTER (WHERE coalesce(trim(cast("完成质量" AS VARCHAR)), '') = ''),
        '情绪', count(*) FILTER (WHERE coalesce(trim(cast("情绪" AS VARCHAR)), '') = ''),
        '执行情境强度', count(*) FILTER (WHERE coalesce(trim(cast("执行情境强度" AS VARCHAR)), '') = ''),
        '状态', count(*) FILTER (WHERE coalesce(trim(cast("状态" AS VARCHAR)), '') = ''),
        '色散类型', count(*) FILTER (WHERE coalesce(trim(cast("色散类型" AS VARCHAR)), '') = ''),
        '认知挑战', count(*) FILTER (WHERE coalesce(trim(cast("认知挑战" AS VARCHAR)), '') = '')
      )
      FROM trace_rows
    ),
    'motion_blank_cells', (
      SELECT
        count(*) FILTER (
          WHERE coalesce(trim("动作"), '') = ''
             OR coalesce(trim("场景"), '') = ''
             OR coalesce(trim("强度"), '') = ''
             OR coalesce(trim("描述"), '') = ''
             OR coalesce(trim("来源"), '') = ''
             OR coalesce(trim("目标区域"), '') = ''
             OR coalesce(trim("英文"), '') = ''
             OR coalesce(trim("象限"), '') = ''
        )
      FROM motion_rows
    ),
    'trace_duplicate_groups', (
      SELECT count(*) FROM trace_duplicate_counts WHERE copies > 1
    ),
    'trace_duplicate_excess_rows', (
      SELECT coalesce(sum(copies - 1), 0)
      FROM trace_duplicate_counts
      WHERE copies > 1
    ),
    'trace_duplicate_groups_with_blank_content', (
      SELECT count(*)
      FROM trace_duplicate_counts
      WHERE copies > 1 AND has_blank_content
    ),
    'trace_nonblank_content_duplicate_keys', (
      SELECT count(*) FROM (
        SELECT "色散内容"
        FROM trace_rows
        WHERE coalesce(trim("色散内容"), '') <> ''
        GROUP BY "色散内容"
        HAVING count(*) > 1
      )
    ),
    'motion_action_duplicate_keys', (
      SELECT count(*) FROM (
        SELECT "动作" FROM motion_rows GROUP BY "动作" HAVING count(*) > 1
      )
    )
  ) AS details

UNION ALL

-- Q3: What time pattern exists in 色散, and how many date values require explicit handling?
SELECT
  'Q3_time_pattern' AS question_id,
  'Monthly distribution plus valid and non-date date-field counts' AS finding,
  json_object(
    'valid_date_rows', (SELECT count(*) FROM trace_dates WHERE trace_date IS NOT NULL),
    'non_date_rows', (SELECT count(*) FROM trace_dates WHERE trace_date IS NULL),
    'first_date', (SELECT min(trace_date) FROM trace_dates),
    'last_date', (SELECT max(trace_date) FROM trace_dates),
    'monthly_counts', (
      SELECT to_json(list(month_row ORDER BY month_row.month))
      FROM (
        SELECT struct_pack(
          month := strftime(trace_date, '%Y-%m'),
          records := count(*)
        ) AS month_row
        FROM trace_dates
        WHERE trace_date IS NOT NULL
        GROUP BY strftime(trace_date, '%Y-%m')
      )
    )
  ) AS details

UNION ALL

-- Q4: Are 探筹's enumerated fields compact, or do they show isolated values/spelling drift?
SELECT
  'Q4_motion_categories' AS question_id,
  'Frequency distributions for scene, intensity, source, and quadrant' AS finding,
  json_object(
    'scene', (
      SELECT to_json(list(struct_pack(value := "场景", records := records) ORDER BY "场景"))
      FROM (SELECT "场景", count(*) AS records FROM motion_rows GROUP BY "场景")
    ),
    'intensity', (
      SELECT to_json(list(struct_pack(value := "强度", records := records) ORDER BY "强度"))
      FROM (SELECT "强度", count(*) AS records FROM motion_rows GROUP BY "强度")
    ),
    'source', (
      SELECT to_json(list(struct_pack(value := "来源", records := records) ORDER BY "来源"))
      FROM (SELECT "来源", count(*) AS records FROM motion_rows GROUP BY "来源")
    ),
    'quadrant', (
      SELECT to_json(list(struct_pack(value := "象限", records := records) ORDER BY "象限"))
      FROM (SELECT "象限", count(*) AS records FROM motion_rows GROUP BY "象限")
    )
  ) AS details

UNION ALL

-- Q5: What schema drift exists across the two CSVs, and can one union-by-name scan preserve both?
SELECT
  'Q5_csv_schema_drift' AS question_id,
  'Inferred schemas are disjoint but union_by_name preserves all rows' AS finding,
  json_object(
    'trace_schema', (
      SELECT to_json(list(struct_pack(name := column_name, type := data_type) ORDER BY column_index))
      FROM duckdb_columns()
      WHERE table_name = 'trace_rows'
    ),
    'motion_schema', (
      SELECT to_json(list(struct_pack(name := column_name, type := data_type) ORDER BY column_index))
      FROM duckdb_columns()
      WHERE table_name = 'motion_rows'
    ),
    'shared_column_count', (
      SELECT count(*)
      FROM (
        SELECT column_name FROM duckdb_columns() WHERE table_name = 'trace_rows'
        INTERSECT
        SELECT column_name FROM duckdb_columns() WHERE table_name = 'motion_rows'
      )
    ),
    'union_counts', (
      SELECT to_json(list(struct_pack(source := source_kind, records := records) ORDER BY source_kind))
      FROM (
        SELECT source_kind, count(*) AS records
        FROM combined_csv_rows
        GROUP BY source_kind
      )
    )
  ) AS details

UNION ALL

-- Q6: Can DuckDB expand the catalog's nested parts and validate its partition contract?
SELECT
  'Q6_json_nested_parts' AS question_id,
  'Nested part objects expand cleanly and can be checked as a sequence' AS finding,
  json_object(
    'catalog_count', (SELECT count FROM catalog_rows),
    'part_rows', (SELECT count(*) FROM catalog_parts),
    'part_count_sum', (SELECT sum(item_count) FROM catalog_parts),
    'catalog_rows_missing_required_fields', (
      SELECT count(*)
      FROM catalog_rows
      WHERE version IS NULL
         OR kind IS NULL
         OR source IS NULL
         OR count IS NULL
         OR part_size IS NULL
         OR fingerprint IS NULL
         OR parts IS NULL
    ),
    'part_rows_missing_required_fields', (
      SELECT count(*)
      FROM catalog_parts
      WHERE path IS NULL
         OR start_id IS NULL
         OR end_id IS NULL
         OR item_count IS NULL
    ),
    'duplicate_paths', (
      SELECT count(*) FROM (
        SELECT path FROM catalog_parts GROUP BY path HAVING count(*) > 1
      )
    ),
    'range_or_count_errors', (
      SELECT count(*)
      FROM catalog_parts
      WHERE end_id - start_id + 1 <> item_count
    ),
    'sequence_gaps', (
      SELECT count(*)
      FROM (
        SELECT start_id, lag(end_id) OVER (ORDER BY start_id) AS previous_end
        FROM catalog_parts
      )
      WHERE previous_end IS NOT NULL AND start_id <> previous_end + 1
    ),
    'parts', (
      SELECT to_json(list(struct_pack(
        path := path,
        start_id := start_id,
        end_id := end_id,
        item_count := item_count
      ) ORDER BY start_id))
      FROM catalog_parts
    )
  ) AS details

UNION ALL

-- Q7 (沃壤真实问题): Is the NLM corpus catalog balanced and internally complete?
SELECT
  'Q7_nlm_partition_health' AS question_id,
  'The catalog can be audited without rebuilding or opening any source card' AS finding,
  json_object(
    'declared_total', (SELECT count FROM catalog_rows),
    'computed_total', (SELECT sum(item_count) FROM catalog_parts),
    'full_parts', (
      SELECT count(*)
      FROM catalog_parts, catalog_rows
      WHERE item_count = part_size
    ),
    'partial_parts', (
      SELECT count(*)
      FROM catalog_parts, catalog_rows
      WHERE item_count < part_size
    ),
    'oversized_parts', (
      SELECT count(*)
      FROM catalog_parts, catalog_rows
      WHERE item_count > part_size
    ),
    'contract_ok', (
      SELECT
        c.count = p.computed_total
        AND p.first_id = 1
        AND p.last_id = c.count
        AND NOT EXISTS (
          SELECT 1
          FROM catalog_parts
          WHERE end_id - start_id + 1 <> item_count
        )
        AND NOT EXISTS (
          SELECT 1
          FROM (
            SELECT start_id, lag(end_id) OVER (ORDER BY start_id) AS previous_end
            FROM catalog_parts
          )
          WHERE previous_end IS NOT NULL AND start_id <> previous_end + 1
        )
      FROM catalog_rows AS c
      CROSS JOIN (
        SELECT
          sum(item_count) AS computed_total,
          min(start_id) AS first_id,
          max(end_id) AS last_id
        FROM catalog_parts
      ) AS p
    )
  ) AS details
ORDER BY question_id;
