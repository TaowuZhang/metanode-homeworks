# 作业1：Solidity 基础题

Remix 可直接打开本目录里的 `.sol` 文件。题目 1~6 各对应一个合约。

| 题 | 文件 | 函数 |
|---|---|---|
| 1 | `Voting.sol` | `vote(uint256)` / `getVotes(uint256)` / `resetVotes()` |
| 2 | `ReverseString.sol` | `reverse(string)` |
| 3 | `IntegerToRoman.sol` | `intToRoman(uint256)`，范围 1..3999 |
| 4 | `RomanToInteger.sol` | `romanToInt(string)` |
| 5 | `MergeSortedArray.sol` | `mergeSorted(uint256[], uint256[])` |
| 6 | `BinarySearch.sol` | `binarySearch(uint256[], uint256)`，找不到返回 `-1` |

作业原文把第 3、4 题的 LeetCode 链接写反了。这里按文字要求实现：第 3 题是整数转罗马，第 4 题是罗马转整数。

## Remix 操作

1. 打开 https://remix.ethereum.org
2. 把对应 `.sol` 复制进 Remix
3. 编译器选 `0.8.28`
4. Environment 选 `Remix VM`
5. Deploy 后按下面用例点按钮

### 1. Voting

- `vote(1)` 两次，`vote(2)` 一次
- `getVotes(1)` 应为 `2`，`getVotes(2)` 应为 `1`，`getVotes(999)` 应为 `0`
- `resetVotes()` 后，`getVotes(1)` 应为 `0`

`resetVotes` 用 epoch 做 O(1) 清零：不遍历 mapping，只把当前轮次加一。旧轮次的票在 `getVotes` 里视为 0。

### 2. ReverseString

- `reverse("abcde")` → `"edcba"`
- `reverse("a")` → `"a"`
- `reverse("")` → `""`

按字节反转，ASCII 字符串没问题。

### 3. IntegerToRoman

- `intToRoman(3749)` → `"MMMDCCXLIX"`
- `intToRoman(1994)` → `"MCMXCIV"`
- `intToRoman(0)` 或 `4000` 会 revert

### 4. RomanToInteger

- `romanToInt("MCMXCIV")` → `1994`
- `romanToInt("IV")` → `4`
- `romanToInt("LVIII")` → `58`

### 5. MergeSortedArray

Remix 数组输入写成 `[1,3,5]`。

- `mergeSorted([1,3,5], [2,4,6])` → `[1,2,3,4,5,6]`
- 任一侧空数组时，结果等于另一侧

### 6. BinarySearch

- `binarySearch([1,3,5,7,9], 5)` → `2`
- `binarySearch([1,3,5,7,9], 4)` → `-1`
- 空数组 → `-1`

## 本地测试

在 `solidity/homework-tests/`：

```bash
cd solidity/homework-tests
npm install
npm test
```
