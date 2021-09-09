package com.teamide.toolbox.util;

import org.springframework.cglib.beans.BeanMap;

import java.lang.reflect.Field;
import java.util.*;

/**
 * 对象和map相互转换工具
 *
 * @author Model Coder
 */
public class BeanMapUtil {

    /**
     * 格式化Map值
     * <p>
     * 属性为long类型，将map中值转为long
     *
     * @param map       mao
     * @param beanClass beanClass
     */
    public static void formatMapValueByBeanClass(Map<String, Object> map, Class<?> beanClass) {
        if (map == null || beanClass == null) {
            return;
        }
        Field[] fields = getAllFields(beanClass);
        for (Field field : fields) {
            String key = field.getName();
            Object value = map.get(key);
            if (value == null) {
                continue;
            }
            if (field.getType().equals(Long.class) || field.getType().equals(long.class)) {
                map.put(key, Long.valueOf(value.toString()));
            } else if (field.getType().equals(Integer.class) || field.getType().equals(int.class)) {
                map.put(key, Integer.valueOf(value.toString()));
            } else if (field.getType().equals(Byte.class)) {
                map.put(key, Byte.valueOf(value.toString()));
            } else if (field.getType().equals(Double.class)) {
                map.put(key, Double.valueOf(value.toString()));
            }
        }
    }

    public static Field[] getAllFields(Class<?> beanClass) {
        List<Field> fieldList = new ArrayList<>();
        while (beanClass != null) {
            fieldList.addAll(new ArrayList<>(Arrays.asList(beanClass.getDeclaredFields())));
            beanClass = beanClass.getSuperclass();
        }
        Field[] fields = new Field[fieldList.size()];
        fieldList.toArray(fields);
        return fields;
    }

    /**
     * 将对象属性转化为map结合
     */
    @SuppressWarnings("unchecked")
    public static Map<String, Object> toMap(Object bean) {
        if (bean instanceof Map) {
            return (Map<String, Object>) bean;
        }
        BeanMap beanMap = BeanMap.create(bean);
        Map<String, Object> map = new HashMap<>(beanMap.size());
        for (Object key : beanMap.keySet()) {
            map.put(String.valueOf(key), beanMap.get(key));
        }
        return map;
    }

    public static List<Map<String, Object>> toMaps(Object beans) {
        return toMaps((List<?>) beans);
    }

    public static List<Map<String, Object>> toMaps(List<?> beans) {
        List<Map<String, Object>> list = new ArrayList<>();
        for (Object bean : beans) {
            Map<String, Object> map = toMap(bean);
            list.add(map);
        }
        return list;
    }

    /**
     * 将map集合中的数据转化为指定对象的同名属性中
     */
    public static <T> T toBean(Map<String, Object> map, Class<T> clazz) throws Exception {
        formatMapValueByBeanClass(map, clazz);
        T bean = clazz.newInstance();
        BeanMap beanMap = BeanMap.create(bean);
        beanMap.putAll(map);
        return bean;
    }

    public static <T> List<T> toBeans(List<Map<String, Object>> maps, Class<T> clazz) throws Exception {
        List<T> list = new ArrayList<>();
        for (Map<String, Object> map : maps) {
            formatMapValueByBeanClass(map, clazz);
            T bean = clazz.newInstance();
            BeanMap beanMap = BeanMap.create(bean);
            beanMap.putAll(map);
            list.add(bean);
        }
        return list;
    }

    public static int getIntValue(Map<String, Object> data, String key) {
        Integer value = getInteger(data, key);
        if (value == null) {
            return 0;
        }
        return value;
    }

    public static Integer getInteger(Map<String, Object> data, String key) {
        Object value = data.get(key);
        if (value == null) {
            return null;
        }
        return Integer.valueOf(value.toString());
    }

    public static byte getByteValue(Map<String, Object> data, String key) {
        Byte value = getByte(data, key);
        if (value == null) {
            return 0;
        }
        return value;
    }

    public static Byte getByte(Map<String, Object> data, String key) {
        Object value = data.get(key);
        if (value == null) {
            return null;
        }
        return Byte.valueOf(value.toString());
    }

    public static boolean getBooleanValue(Map<String, Object> data, String key) {
        Boolean value = getBoolean(data, key);
        if (value == null) {
            return false;
        }
        return value;
    }

    public static Boolean getBoolean(Map<String, Object> data, String key) {
        Object value = data.get(key);
        if (value == null) {
            return null;
        }
        return Boolean.valueOf(value.toString());
    }

    public static String getString(Map<String, Object> data, String key) {
        Object value = data.get(key);
        if (value == null) {
            return null;
        }
        return value.toString();
    }
}