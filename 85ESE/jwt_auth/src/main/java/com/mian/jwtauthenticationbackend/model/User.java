package com.mian.jwtauthenticationbackend.model;

import io.swagger.v3.oas.annotations.media.Schema;
import jakarta.persistence.*;
import org.springframework.security.core.GrantedAuthority;
import org.springframework.security.core.authority.SimpleGrantedAuthority;
import org.springframework.security.core.userdetails.UserDetails;

import java.util.Collection;
import java.util.List;

@Entity
@Table(name="users")
@Schema(description = "Usuário do sistema")
public class User implements UserDetails {

    @Id
    @GeneratedValue (strategy = GenerationType.IDENTITY)
    @Column(name="id")
    @Schema(description = "ID do usuário", example = "1")
    private Integer id;

    @Column(name="first_name")
    @Schema(description = "Nome", example = "João")
    private String firstName;
    @Column(name="last_name")
    @Schema(description = "Sobrenome", example = "Silva")
    private String lastname;

    @Column(name="password")
    @Schema(description = "Senha", example = "senha123")
    private String password;

    @Column(name="username")
    @Schema(description = "Nome de usuário", example = "joao")
    private String username;
    @Enumerated(value= EnumType.STRING)
    @Schema(description = "Papel do usuário", example = "ADMIN")
    private Role role;

    public void setId(Integer id) {
        this.id = id;
    }

    public void setFirstName(String firstName) {
        this.firstName = firstName;
    }

    public void setLastname(String lastname) {
        this.lastname = lastname;
    }

    public void setPassword(String password) {
        this.password = password;
    }

    public void setUsername(String username) {
        this.username = username;
    }

    public void setRole(Role role) {
        this.role = role;
    }

    public Integer getId() {
        return id;
    }

    public String getFirstName() {
        return firstName;
    }

    public String getLastname() {
        return lastname;
    }
    public String getPassword() {
        return password;
    }

    public String getUsername() {
        return username;
    }
    @Override
    public Collection<? extends GrantedAuthority> getAuthorities() {
        return List.of(new SimpleGrantedAuthority(role.name()));
    }


    @Override
    public boolean isAccountNonExpired() {
        return true;
    }

    @Override
    public boolean isAccountNonLocked() {
        return true;
    }

    @Override
    public boolean isCredentialsNonExpired() {
        return true;
    }

    @Override
    public boolean isEnabled() {
        return true;
    }

    public Role getRole() {
        return role;
    }
}