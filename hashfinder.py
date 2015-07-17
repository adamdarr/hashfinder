import hashlib
import sys

goal = ('f04031af5ee9939a931a062b836c385eb2ee681fd704ed919534b78329ae34e173'
        'e13c31ebfd03db1eb9e9759cd3365195d0d7d53028256038b24975ace8c83d')

for i in range(256):
    for j in range(256):
        ip = "172.16.%d.%d" % (i, j)
        digest = hashlib.sha512(ip).hexdigest()
        if digest == goal:
            print ip
            sys.exit(0)