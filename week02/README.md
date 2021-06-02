## 题目

在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？

## 思路

1. 分析题目
2. 模拟出sql.ErrNoRows的错误
3. 处理sql.ErrNoRows错误
4. 测试代码和整理代码
