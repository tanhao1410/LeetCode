from typing import List


class Solution:
    # 22. 括号生成
    def generateParenthesis(self, n: int) -> List[str]:
        # n个括号，说名字符串长度是2n,生成字符串时，每一次可以选择的有左或右。最多是2**16个
        return self.createStr('', 0, n)

    def createStr(self, pre: str, left: int, n: int) -> List[str]:
        res = []
        if left == n:
            # 左括号的数量达到了最大，则剩下的只能是右括号
            res.append(pre + ')' * (2 * n - len(pre)))
        else:
            # 两种情况，可以选择左，也可以选择右括号，
            res += self.createStr(pre + '(', left + 1, n)
            # 右括号不能放在前面
            if len(pre) - left < left:
                res += self.createStr(pre + ')', left, n)
        return res

    # 824. 山羊拉丁文
    def toGoatLatin(self, sentence: str) -> str:
        letters = {'a', 'e', 'i', 'o', 'u'}
        words = sentence.split(" ")
        res = []
        for index, word in enumerate(words):
            if word[0].lower() not in letters:
                word = word[1:] + word[0]
            word += 'ma'
            word += 'a' * (index + 1)
            res.append(word)
        return ' '.join(res)

    # 88. 合并两个有序数组
    def merge(self, nums1: List[int], m: int, nums2: List[int], n: int) -> None:
        """
        Do not return anything, modify nums1 in-place instead.
        """
        p1, p2, p3 = m - 1, n - 1, m + n - 1
        while p3 >= 0:
            if p2 < 0 or p1 >= 0 and nums1[p1] > nums2[p2]:
                nums1[p3] = nums1[p1]
                p1 -= 1
            else:
                nums1[p3] = nums2[p2]
                p2 -= 1
            p3 -= 1
