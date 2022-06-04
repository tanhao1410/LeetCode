from typing import List


class Solution:
    # 929. 独特的电子邮件地址
    def numUniqueEmails(self, emails: List[str]) -> int:
        email_set = set()
        for email in emails:
            name_host = email.split('@')
            name = name_host[0].replace('.', '').split('+')[0]
            # 如果有.则消除
            email_set.add(name + '@' + name_host[1])
        return len(email_set)
