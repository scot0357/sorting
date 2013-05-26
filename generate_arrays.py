import random

STRING_TMPL = "arraysize-10e%d-test%d.dat"

def generate_file(f, size):
    for x in xrange(size):
        x = random.randrange(-200000,200000)
        f.write(str(x) + "\n")

def main():
    for x in range(1,7):
        for y in range(1,4):
            filename = STRING_TMPL % (x, y)
            fobject = open(filename, "wb")
            generate_file(fobject, 10**x)

if __name__ == "__main__":
    main()
