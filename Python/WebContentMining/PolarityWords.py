import urllib
import os.path

files=['negative.txt','positive.txt','obama_tweets.txt']
path='http://www.unc.edu/~ncaren/haphazard/'
for file_name in files:
	if os.path.isfile(file_name) != True:      #if file doesn't exist
		urllib.urlretrieve(path+file_name,file_name)
