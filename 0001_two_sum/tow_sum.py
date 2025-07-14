def two_sum(nums, target):
    num_map = {}  # 用来存储数值和对应的下标
    for i, num in enumerate(nums):
        complement = target - num  # 计算差值
        if complement in num_map:
            return [num_map[complement], i]  # 找到配对
        num_map[num] = i  # 没找到就记录当前数值和下标
    return []

# 示例：
nums = [2, 7, 11, 15]
target = 9
result = two_sum(nums, target)
print(result)  # 输出 [0, 1]

