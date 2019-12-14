package com.github.MickiFoerster.protobuf;

import com.example.Simple.SimpleMessage;
import com.google.protobuf.*;
import com.google.protobuf.util.JsonFormat;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.FileOutputStream;
import java.io.IOException;
import java.util.Arrays;

public class ProtoToJSONMain {
  public static void main(String[] args) {
    System.out.println("Hello from ProtoToJSON main!");

    SimpleMessage.Builder builder = SimpleMessage.newBuilder();

    // simple field
    builder.setId(42);
    builder.setIsSimple(true);
    builder.setName("This is a simpel message");

    builder.setId(23).setIsSimple(true).setName("This is another message");

    // repeated field
    builder.addSampleList(1).addSampleList(2).addSampleList(3).addSampleList(4);

    builder.addAllSampleList(Arrays.asList(4, 5, 6));

    // Print this as a JSON
    String jsonString = new String();
    try {
      jsonString = JsonFormat.printer().print(builder);
      System.out.println(jsonString);
    } catch (InvalidProtocolBufferException e) {
      e.printStackTrace();
    }

    // Parse JSON into protobuf
    SimpleMessage.Builder builder2 = SimpleMessage.newBuilder();
    try {
      JsonFormat.parser().ignoringUnknownFields().merge(jsonString, builder2);
    } catch (InvalidProtocolBufferException e) {
      e.printStackTrace();
    }
    System.out.println(builder2);
  }
}