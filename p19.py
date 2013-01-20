import datetime

start = datetime.date(1901, 1, 1)
end = datetime.date(2001, 1, 1)

d = start
t = datetime.timedelta(days=1)

cnt = 0
while d < end:
    if d.weekday() == 6 and d.day == 1:
        cnt += 1
    d += t

print(cnt)
