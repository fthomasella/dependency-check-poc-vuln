package com.example;

import org.apache.commons.collections.map.HashedMap;

public class App {
    public static void main(String[] args) {
        HashedMap map = new HashedMap();
        map.put("key", "value");
        System.out.println("PoC para Dependency-Check: " + map.get("key"));
    }
}
