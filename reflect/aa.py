


def foo(a):
    print(a)


if __name__ == '__main__':
    aa = __import__("aa")

    f = getattr(aa,"foo")
    print(f)

    f("6")



    