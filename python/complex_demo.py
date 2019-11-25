import complex.complex_pb2 as complex_pb2

complex_message = complex_pb2.ComplexMessage()
complex_message.one_dummy.id = 1234
complex_message.one_dummy.name = "I am a dummy message"

# first possibility
first_multiple_dummy = complex_message.multiple_dummy.add()
first_multiple_dummy.id = 345
first_multiple_dummy.name = "I'm the first element of multiple dummies"

# second possibility
complex_message.multiple_dummy.add(
    id=567, name="2nd element of multiple dummies")

# third possibility
third_dummy_message = complex_pb2.DummyMessage()
third_dummy_message.id = 999
third_dummy_message.name = "I'm the last dummy"

complex_message.multiple_dummy.extend([third_dummy_message])
third_dummy_message.id = 666

print(complex_message)
