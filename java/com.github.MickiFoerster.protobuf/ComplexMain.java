package com.github.MickiFoerster.protobuf;

import java.util.Arrays;
import java.io.FileOutputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

import example.complex.Complex;
import example.complex.Complex.ComplexMessage;

public class ComplexMain {
    public static void main(String[] args) {
        System.out.println("Example for Complex Messages");

        Complex.DummyMessage dummyMsg = newDummyMessage(55, "one dummy message");

        ComplexMessage.Builder builder = ComplexMessage.newBuilder();

        builder.setOneDummy(dummyMsg);

        // repeated field
        builder.addMultipleDummy(newDummyMessage(66, "second dummy"));
        builder.addMultipleDummy(newDummyMessage(67, "third dummy"));
        builder.addMultipleDummy(newDummyMessage(68, "fourth dummy"));

        ComplexMessage msg = builder.build();

        System.out.println(msg.toString());

        // message.getMultipleDummyList();
    }

    public static Complex.DummyMessage newDummyMessage(Integer id, String name) {
        Complex.DummyMessage.Builder dummyMessageBuilder = Complex.DummyMessage.newBuilder();
        return dummyMessageBuilder.setName(name).setId(id).build();
    }
}
