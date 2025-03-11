package com.example.blog.controller;

import com.example.blog.model.Post;
import com.example.blog.model.User;
import com.example.blog.service.PostService;
import com.example.blog.service.UserService;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.security.core.Authentication;
import org.springframework.security.core.context.SecurityContextHolder;
import org.springframework.stereotype.Controller;
import org.springframework.ui.Model;
import org.springframework.web.bind.annotation.*;
import org.springframework.web.servlet.mvc.support.RedirectAttributes;

import java.time.LocalDateTime;

@Controller
@RequestMapping("/posts")
public class PostController {

    private final PostService postService;
    private final UserService userService;

    @Autowired
    public PostController(PostService postService, UserService userService) {
        this.postService = postService;
        this.userService = userService;
    }

    @GetMapping("/new")
    public String showNewPostForm(Model model) {
        model.addAttribute("post", new Post());
        return "post-form";
    }

    @PostMapping("/new")
    public String createPost(@ModelAttribute Post post, RedirectAttributes redirectAttributes) {
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        User currentUser = userService.findUserByUsername(auth.getName()).orElseThrow();
        
        post.setAuthor(currentUser);
        post.setCreatedAt(LocalDateTime.now());
        
        postService.createPost(post);
        redirectAttributes.addFlashAttribute("success", "文章发布成功");
        return "redirect:/";
    }

    @GetMapping("/{id}")
    public String showPost(@PathVariable Long id, Model model) {
        Post post = postService.findPostById(id).orElseThrow();
        model.addAttribute("post", post);
        return "post-detail";
    }

    @GetMapping("/{id}/edit")
    public String showEditForm(@PathVariable Long id, Model model) {
        Post post = postService.findPostById(id).orElseThrow();
        
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        User currentUser = userService.findUserByUsername(auth.getName()).orElseThrow();
        
        if (!post.getAuthor().getId().equals(currentUser.getId())) {
            return "redirect:/";
        }
        
        model.addAttribute("post", post);
        return "post-form";
    }

    @PostMapping("/{id}/edit")
    public String updatePost(@PathVariable Long id, @ModelAttribute Post post, RedirectAttributes redirectAttributes) {
        Post existingPost = postService.findPostById(id).orElseThrow();
        
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        User currentUser = userService.findUserByUsername(auth.getName()).orElseThrow();
        
        if (!existingPost.getAuthor().getId().equals(currentUser.getId())) {
            return "redirect:/";
        }
        
        existingPost.setTitle(post.getTitle());
        existingPost.setContent(post.getContent());
        existingPost.setUpdatedAt(LocalDateTime.now());
        
        postService.updatePost(existingPost);
        redirectAttributes.addFlashAttribute("success", "文章更新成功");
        return "redirect:/posts/" + id;
    }

    @PostMapping("/{id}/delete")
    public String deletePost(@PathVariable Long id, RedirectAttributes redirectAttributes) {
        Post post = postService.findPostById(id).orElseThrow();
        
        Authentication auth = SecurityContextHolder.getContext().getAuthentication();
        User currentUser = userService.findUserByUsername(auth.getName()).orElseThrow();
        
        if (!post.getAuthor().getId().equals(currentUser.getId())) {
            return "redirect:/";
        }
        
        postService.deletePost(id);
        redirectAttributes.addFlashAttribute("success", "文章删除成功");
        return "redirect:/";
    }
} 