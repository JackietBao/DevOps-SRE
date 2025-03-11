package com.example.blog.service;

import com.example.blog.model.Post;
import com.example.blog.model.User;

import java.util.List;
import java.util.Optional;

public interface PostService {
    Post createPost(Post post);
    Post updatePost(Post post);
    void deletePost(Long id);
    Optional<Post> findPostById(Long id);
    List<Post> findAllPosts();
    List<Post> findPostsByAuthor(User author);
} 