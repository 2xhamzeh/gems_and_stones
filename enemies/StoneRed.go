components {
  id: "stone"
  component: "/enemies/enemy.script"
  properties {
    id: "speed"
    value: "30.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "health"
    value: "50.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "hit_value"
    value: "20.0"
    type: PROPERTY_TYPE_NUMBER
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"stone_red\"\n"
  "material: \"/builtins/materials/sprite.material\"\n"
  "textures {\n"
  "  sampler: \"texture_sampler\"\n"
  "  texture: \"/assets/pngs.atlas\"\n"
  "}\n"
  ""
}
embedded_components {
  id: "collisionobject"
  type: "collisionobject"
  data: "type: COLLISION_OBJECT_TYPE_TRIGGER\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"stone\"\n"
  "mask: \"player\"\n"
  "mask: \"attack\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 7.524085\n"
  "}\n"
  ""
}
