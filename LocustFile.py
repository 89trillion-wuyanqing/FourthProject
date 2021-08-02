
from locust import HttpUser, TaskSet, task, between


class UserBehavior(TaskSet):

    @task(1)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def registerUser(self):

        data = {
            "id": "100001"

        }
        response = self.client.post('/createGiftCodes', data = data,name="registerUser")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')

    @task(2)  # 任务的权重为1，如果有多个任务，可以将权重值定义成不同的值，
    def activateCode(self):

        data = {
            "giftCode": "25EXYA5U",
            "userId":"100001"

        }
        response = self.client.post('/activateCode', data = data,name="activateCode")
        if not response.ok:
            print(response.text)
            response.failure('Got wrong response')



class TestLocust(HttpUser):

    tasks = [UserBehavior]
    wait_time = between(2, 5)
    host = "http://127.0.0.1:8000"
    #task_set = UserBehavior
    #host = "http://127.0.0.1/:8000"  # 被测服务器地址
    #min_wait = 5000
# 最小等待时间，即至少等待多少秒后Locust选择执行一个任务。

    #max_wait = 9000
# 最大等待时间，即至多等待多少秒后Locust选择执行一个任务。