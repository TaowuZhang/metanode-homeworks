#!/usr/bin/env bash
set -euo pipefail

usage() {
  echo "usage: $0 TRACE_CSV MOTION_CSV CATALOG_JSON [json|csv]" >&2
}

if [[ $# -lt 3 || $# -gt 4 ]]; then
  usage
  exit 64
fi

trace_input=$1
motion_input=$2
catalog_input=$3
output_format=${4:-json}
duckdb_bin=${DUCKDB_BIN:-duckdb}

for input in "$trace_input" "$motion_input" "$catalog_input"; do
  if [[ ! -f "$input" ]]; then
    echo "input is not a regular file: $input" >&2
    exit 66
  fi
done

case "$output_format" in
  json) output_flag=-json ;;
  csv) output_flag=-csv ;;
  *)
    echo "unsupported output format: $output_format" >&2
    usage
    exit 64
    ;;
esac

if ! command -v "$duckdb_bin" >/dev/null 2>&1; then
  echo "DuckDB CLI not found: $duckdb_bin" >&2
  exit 69
fi

script_dir=$(cd -- "$(dirname -- "${BASH_SOURCE[0]}")" && pwd -P)
sql_file="$script_dir/duckdb-local-lens-experiment.sql"
if [[ ! -f "$sql_file" ]]; then
  echo "experiment SQL not found: $sql_file" >&2
  exit 66
fi

temp_dir=$(mktemp -d "${TMPDIR:-/tmp}/wo-duckdb-local-lens.XXXXXX")
cleanup() {
  rm -rf "$temp_dir"
}
trap cleanup EXIT HUP INT TERM

# Copies keep every query inside a disposable directory and exercise Chinese paths with spaces.
trace_copy="$temp_dir/色散 样本.csv"
motion_copy="$temp_dir/探筹 样本.csv"
catalog_copy="$temp_dir/NLM 目录 样本.json"
cp "$trace_input" "$trace_copy"
cp "$motion_input" "$motion_copy"
cp "$catalog_input" "$catalog_copy"

sql_quote() {
  sed "s/'/''/g" <<<"$1"
}

trace_sql=$(sql_quote "$trace_copy")
motion_sql=$(sql_quote "$motion_copy")
catalog_sql=$(sql_quote "$catalog_copy")
result_file="$temp_dir/result.$output_format"

"$duckdb_bin" :memory: "$output_flag" \
  -cmd "SET VARIABLE trace_path = '$trace_sql';" \
  -cmd "SET VARIABLE motion_path = '$motion_sql';" \
  -cmd "SET VARIABLE catalog_path = '$catalog_sql';" \
  <"$sql_file" >"$result_file"

if [[ "$output_format" == json ]]; then
  jq -e 'type == "array" and length >= 7' "$result_file" >/dev/null
fi

cat "$result_file"
