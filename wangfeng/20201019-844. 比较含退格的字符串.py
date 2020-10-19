class Solution:
    def backspaceCompare(self, S: str, T: str) -> bool:

        stack1 = []
        stack2 = []
        for c in S:
            if c == '#':
                if len(stack1) != 0:
                    stack1.pop()
            else:
                stack1.append(c)
        for c in T:
            if c == '#':
                if len(stack2) != 0:
                    stack2.pop()
            else:
                stack2.append(c)
        
        return stack1==stack2
