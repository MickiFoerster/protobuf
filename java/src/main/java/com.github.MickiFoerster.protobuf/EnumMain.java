package com.github.MickiFoerster.protobuf;

import java.util.Arrays;
import java.io.FileOutputStream;
import java.io.FileInputStream;
import java.io.FileNotFoundException;
import java.io.IOException;

import example.enumerations.EnumExample;
import example.enumerations.EnumExample.EnumMessage;

public class EnumMain {
    public static void main(String[] args) {
        System.out.println("Example for Enums");

        EnumMessage.Builder builder = EnumMessage.newBuilder();
        builder.setId(345);
        builder.setDayOfTheWeek(EnumExample.DayOfTheWeek.FRIDAY);

        EnumMessage message = builder.build();
        System.out.println(message);
    }
}
