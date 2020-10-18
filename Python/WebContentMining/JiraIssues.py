from jira.client import JIRA
from datetime import datetime
import Report
import csv
import Sendmail

class Report(object):
    def __init__(self, project_name=None, task_number=None, task_type=None, task_summary=None, task_duedate=None, assignee=None,createDate=None):
        self.project_name = project_name
        self.task_number = task_number
        self.task_type = task_type
        self.task_summary = task_summary
        self.task_duedate = task_duedate
        self.assignee = assignee
        self.createDate = createDate
class Ont:
reportNames=["reportOne.csv","reportTwo.csv"]
    def connect_jira(self,jira_server, jira_user, jira_password):
        '''
        Connect to JIRA. Return None on error
        '''
        try:
            jira_options = {'server': jira_server}
            jira = JIRA(options=jira_options, basic_auth=(jira_user, jira_password))
        except Exception as e:
            print e
            jira = None
        return jira
ont = Ont()
jira = ont.connect_jira("https://foo.jira.com", "fooName", "fooPass")
projects = jira.projects()

def days_between(createdate, todaydate):
    createdate = datetime.strptime(createdate[0:10], "%Y-%m-%d")
    todaydate = datetime.strptime(str(todaydate)[0:10], "%Y-%m-%d")
    return abs((todaydate - createdate).days)

def writeReport(jiraReport,reportName):
    with open(reportName,'wb') as f:
        writer = csv.writer(f)
        writer.writerows(jiraReport)

def listReport():
    '''
    Return List of all Issues
    '''
    try:
        reportList= []
        keys = [project.key for project in projects]
        for key in keys :
            try:
                proj_key = jira.project(key)
                issues = jira.search_issues('project='+key)
                for issue in issues:
                    isstype=issue.fields.issuetype
                    summary=issue.fields.summary
                    due=issue.fields.duedate
                    ass=issue.fields.assignee
                    createdate = issue.fields.created
                    reportList.append(Report.Report(proj_key.name,issue,isstype,summary,due,ass,createdate))
            except Exception as Ex:
                pass
                continue
    except Exception as ex:
        pass
    return reportList

reports=listReport()
reportOne=[]
reportTwo=[]
todaydate = datetime.now()

for report in reports:
    day = days_between(report.createDate, todaydate)
    report_values=[str(report.project_name.encode('utf8')),str(report.task_number),str(report.task_type),str(report.task_summary.encode('utf8')),str(report.task_duedate),str(report.assignee)]
    if (day <= 90):
        reportTwo.append(report_values)
    reportOne.append(report_values)

writeReport(reportOne,reportNames[0])
writeReport(reportTwo,reportNames[1])


