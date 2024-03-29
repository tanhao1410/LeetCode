from typing import List


class Solution:
    # 2000. 反转单词前缀
    def reversePrefix(self, word: str, ch: str) -> str:
        index = word.find(ch)
        if index > 0:
            res = ch
            for i in range(index - 1, -1, -1):
                res += word[i]
            return res + word[index + 1:]
        return word

    # 1037. 有效的回旋镖
    def isBoomerang(self, points: List[List[int]]) -> bool:
        return not (points[1][1] - points[0][1]) * (points[2][0] - points[0][0]) == (points[1][0] - points[0][0]) * (
                points[2][1] - points[0][1])

    # 812. 最大三角形面积
    def largestTriangleArea(self, points: List[List[int]]) -> float:
        def triangleArea(p1, p2, p3) -> float:
            return 0.5 * abs(
                p1[0] * p2[1] + p2[0] * p3[1] + p3[0] * p1[1] - p1[1] * p2[0] - p2[1] * p3[0] - p3[1] * p1[0])

        res = 0
        for i in range(len(points) - 2):
            for j in range(i + 1, len(points) - 1):
                for k in range(j + 1, len(points)):
                    res = max(res, triangleArea(points[i], points[j], points[k]))
        return res

    # 908. 最小差值 I
    def smallestRangeI(self, nums: List[int], k: int) -> int:
        # return max(0,max(nums) - min(nums) - 2 * k)
        max_, min_ = max(nums), min(nums)
        if max_ - min_ > 2 * k:
            return max_ - min_ - 2 * k
        return 0
