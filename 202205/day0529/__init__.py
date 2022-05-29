class Solution:
    # 468. 验证IP地址
    def validIPAddress(self, queryIP: str) -> str:

        # 先判断是否是v4
        if '.' in queryIP:
            ips = queryIP.split('.')
            if all(map(self.is_ip4_num, iter(ips))) and len(ips) == 4:
                return 'IPv4'
        elif ':' in queryIP:
            ips = queryIP.split(':')
            if all(map(self.is_ip6_num, iter(ips))) and len(ips) == 8:
                return 'IPv6'
        return 'Neither'

    def is_ip4_num(self, num: str) -> bool:
        nums = '0123456789'
        if len(num) < 1:
            return False
        if len(num) > 1 and num[0] == '0':
            return False
        # 每一位都应该是数字
        for bit in num:
            if bit not in nums:
                return False
        return 0 <= int(num) <= 255

    def is_ip6_num(self, num: str) -> bool:
        nums = '0123456789abcdefABCDEF'
        if 1 <= len(num) <= 4:
            for bit in num:
                if bit not in nums:
                    return False
            return True
        return False
