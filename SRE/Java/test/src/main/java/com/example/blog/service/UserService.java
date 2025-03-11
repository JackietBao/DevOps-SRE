package com.example.blog.service;

import com.example.blog.model.User;

import java.util.List;
import java.util.Optional;

public interface UserService {
    User registerNewUser(User user);
    List<User> findAllUsers();
    Optional<User> findUserById(Long id);
    Optional<User> findUserByUsername(String username);
    boolean existsByUsername(String username);
    boolean existsByEmail(String email);
} 