from itertools import groupby
step = "1113122113"
for i in range(50):
    p = [list(g) for k, g in groupby(step)]
    step = "".join(["%d%s" % (len(l), l[0]) for l in p])
    if i == 39:
        print(len(step))

print(len(step))
