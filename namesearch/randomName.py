import random

result = []

for i in range(2):
    for j in range(3):
        key1 = random.randint(1, 6500)
        key2 = random.randint(1, 3500)
        result.append((key1, key2))

for k in result:
    print(k[0], k[1])

