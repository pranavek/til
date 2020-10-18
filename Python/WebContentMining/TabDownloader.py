___author___ = "@pranavek"

import wget
from urllib2 import urlopen
from lxml.html import fromstring

def get_page(url):
	html = urlopen(url).read()
	dom = fromstring(html)
	dom.make_links_absolute(url)
	return dom

def download_from_scene(filePageUrl):
	filePage = get_page(filePageUrl)
	downloadLink = filePage.cssselect("#main_scene_box > div:nth-child(1) > div:nth-child(3) > div > a")
        wget.download(downloadLink[0].attrib['href'])
	
for page in range(1,601):
	dom = get_page('http://m.foo.com/videos?page=%d'% page)
	links =	dom.cssselect("div.details > h2 > a")
	for link in links:
		download_from_scene(link.attrib['href'])



