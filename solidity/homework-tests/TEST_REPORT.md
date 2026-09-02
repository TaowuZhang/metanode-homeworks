# Test Report — homework01 / homework02

Date: 2026-09-02

## Commands

```bash
cd solidity/homework-tests
npm install
npm test
```

## Result

- Hardhat Mocha: **11 passing / 0 failing**

### homework01

- Voting: 投票累计、查询未投票候选人为 0、reset 后清零并可重新投票
- ReverseString: `"abcde"` → `"edcba"`，单字符和空串
- IntegerToRoman: `3749` / `1994` / `58`，越界 revert
- RomanToInteger: `"MCMXCIV"` / `"IV"` / `"LVIII"`，非法字符 revert
- MergeSortedArray: 两侧有值、一侧为空
- BinarySearch: 命中首/中/尾、未找到、空数组

### homework02

- donate + owner withdraw（`address.transfer`）
- MetaMask 式普通转账走 `receive()`
- 零值捐赠、非 owner 提取、空余额提取
- Top3 累计金额排行
- 时间窗外捐赠被拒绝，进入窗口后可捐
