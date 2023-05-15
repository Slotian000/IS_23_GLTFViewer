#version 410 core

in vec2 BaseCoord;

out vec4 FragColor;

uniform sampler2D base;

void main()
{
    FragColor = texture(base, BaseCoord);
}