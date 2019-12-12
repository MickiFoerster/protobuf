package com.github.MickiFoerster.protobuf;

import com.example.Simple.SimpleMessage;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class SimpleMain {
  public static void main(String[] args) {
    System.out.println("Hello World!");

    SimpleMessage.Builder builder = SimpleMessage.newBuilder();

    // simple field
    builder.setId(42);
    builder.setIsSimple(true);
    builder.setName("This is a simpel message");

    builder.setId(23).setIsSimple(true).setName("This is another message");

    // repeated field
    builder.addSampleList(1).addSampleList(2).addSampleList(3).addSampleList(4);

    builder.addAllSampleList(Arrays.asList(4, 5, 6));

    builder.setSampleList(3, 42);

    System.out.println(builder.toString());

    builder.clearSampleList();

    // write message to file
    SimpleMessage message = builder.build();
    try {
      FileOutputStream outputStream =
          new FileOutputStream("simple_message.bin");
      message.writeTo(outputStream);
      outputStream.close();
    } catch (FileNotFoundException e) {
      e.printStackTrace();
    } catch (IOException e) {
      e.printStackTrace();
    }

    // sent as byte array
    // byte[] bytes = message.toByteArray();

    try {
      System.out.println("Reading from file ...");
      FileInputStream fileInputStream =
          new FileInputStream("simple_message.bin");
      SimpleMessage messageFromFile = SimpleMessage.parseFrom(fileInputStream);
      System.out.println(messageFromFile);
    } catch (FileNotFoundException e) {
      e.printStackTrace();
    } catch (IOException e) {
      e.printStackTrace();
    }
  }
}
