class Solution:
    def maxProfit(self, prices: List[int]) -> int:
        i = 0
        size = len(prices)
        if size <= 1:
            return 0
        profit = 0
        balance = -1
        while i < size-1:
            if balance != -1:
                if prices[i] > prices[i+1]:
                    profit += prices[i] - balance
                    balance = -1
            else:
                if prices[i] < prices[i+1]:
                    balance = prices[i]
            i += 1
        if balance != -1:
            profit += prices[size-1] - balance
        return profit