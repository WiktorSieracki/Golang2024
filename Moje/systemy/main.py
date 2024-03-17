import random 

with open('pogrubione.txt', 'r',encoding='utf-8') as file:
    tab = file.readlines()

tab = [x.strip() for x in tab]


while True:
    # print(random.choice(tab))
    sample = random.sample(tab, 10)
    for i in sample:
        print(i)
    input()