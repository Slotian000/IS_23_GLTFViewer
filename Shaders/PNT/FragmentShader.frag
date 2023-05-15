#version 410 core

in vec2 TextureCoord;

out vec4 FragColor;

uniform sampler2D base;
//uniform sampler2D normalMap;

void main()
{
    FragColor = texture(base, TextureCoord);
    //FragColor = vec4(1,0,0,1);
}