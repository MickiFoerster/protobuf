import simple.simple_pb2 as simple_pb2

sm = simple_pb2.SimpleMessage()

sm.id = 23
sm.is_simple = True
sm.name = "This is the name of the message"
sml = sm.sample_list
sml.append(3)
sml.append(4)
sml.append(5)

print("created message:")
print(sm)

with open("simple.bin", "wb") as f:
    bytesAsString = sm.SerializeToString()
    f.write(bytesAsString)
    print('message was written to file')
    f.close()

with open("simple.bin", "rb") as f:
    sm_msg_read = simple_pb2.SimpleMessage().FromString(f.read())
    f.close()
    print(sm_msg_read)
    print("Only one value:", sm_msg_read.is_simple)
    print("Only one value:", sm_msg_read.name)
