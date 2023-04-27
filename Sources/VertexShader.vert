#version 410 core

layout (location =0) in vec3 position;
layout (location =1) in vec3 normal;
layout (location =2) in vec2 baseCoord;
layout (location =3) in vec2 normalCoord;

//out vec3 color;
out vec2 BaseCoord;
out vec2 NormalCoord;

uniform mat4 model;
uniform mat4 view;
uniform mat4 projection;



void main()
{
    gl_Position = projection * view * model * vec4(position, 1.0f);
    BaseCoord = baseCoord;
    NormalCoord = normalCoord;
}