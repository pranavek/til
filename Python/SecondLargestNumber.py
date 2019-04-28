
# coding: utf-8

# In[1]:

input_array = [2,3,5,8,1,9]


# In[3]:

first = second = 0
for i in input_array[:]:
    if i > first:
        second = first
        first = i
    elif i > second and i != first:
        second = i


# In[4]:

print("The largest number is {} and the second largest number is {}".format(first, second))


# In[ ]:



