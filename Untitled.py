#!/usr/bin/env python
# coding: utf-8

# In[13]:


S = input()

if len(S) == 0:
    max_length = 0
elif len(S) == 1:
    max_length = 1
else:
    max_length = 1
    current_length = 1
    for i in range(1, len(S)):
        if S[i] == S[i-1]:
            current_length += 1
            if current_length > max_length:
                max_length = max(max_length, current_length)
        else:
            current_length = 1
print(max_length)

