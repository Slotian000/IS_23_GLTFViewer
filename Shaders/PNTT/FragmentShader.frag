#version 410 core

in vec2 BaseCoord;
//in vec2 NormalCoord;
//in vec3 temp;

out vec4 FragColor;

uniform sampler2D base;
//uniform sampler2D normalMap;


void main()
{
    FragColor = texture(base, BaseCoord);
    //FragColor = vec4(1,0,0,1);
    //FragColor = vec4(temp,1);
}