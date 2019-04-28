#Array rotation using temp array
input_array = [1,2,3,4,5,6,7,8,9,0]
rotate_index = 1

def rotate_array(input_array,rotate_index):
    if rotate_index > len(input_array):
        raise Exception("Rotate index cannot be more than input array")
        
    start_index = len(input_array) -  rotate_index
    
    temp_array = []
    for i in range(start_index,len(input_array)):
        temp_array.append(input_array[i])
    for i in range(0,start_index):
        temp_array.append(input_array[i])
    print(temp_array)
     
rotate_array(input_array, rotate_index)    
