import enums.enum_example_pb2 as enum_example_pb2

enum_message = enum_example_pb2.EnumMessage()

enum_message.id = 345
enum_message.day_of_the_week = enum_example_pb2.THURSDAY

print(enum_message)

with open("enums.bin", "wb") as f:
    f.write(enum_message.SerializeToString())
    f.close()


with open("enums.bin", "rb") as f:
    msg = enum_example_pb2.EnumMessage().FromString(f.read())
    f.close()
    print("read from file:")
    print(msg)
