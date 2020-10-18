import requests
import time
import re

root_url = 'http://###.##.###.##'
index_url = root_url + '/CuPbhavan/res_newregentry.php?id='
regexp = re.compile(r'B.Tech \(09 Scheme\)')
f = open('btech_result.09', 'a')

def get_result_page_urls():
	for id in range(4000,5000):
		url = index_url+str(id)
		response = requests.get(url)
		if regexp.search(response.text) is not None:
			f.write(url)
			f.flush()
		time.sleep(1)
		
get_result_page_urls()
f.close()

