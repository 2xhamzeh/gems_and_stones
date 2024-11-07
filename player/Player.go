components {
  id: "player"
  component: "/player/player.script"
}
components {
  id: "attack_factory"
  component: "/player/attacks/attack_factory.script"
  properties {
    id: "factory_id"
    value: "attack"
    type: PROPERTY_TYPE_URL
  }
}
components {
  id: "ability_factory"
  component: "/player/attacks/attack_factory.script"
  properties {
    id: "attack_rate"
    value: "1.0"
    type: PROPERTY_TYPE_NUMBER
  }
  properties {
    id: "input_trigger"
    value: "ability"
    type: PROPERTY_TYPE_HASH
  }
  properties {
    id: "factory_id"
    value: "ability"
    type: PROPERTY_TYPE_URL
  }
}
embedded_components {
  id: "sprite"
  type: "sprite"
  data: "default_animation: \"diamond\"\n"
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
  data: "type: COLLISION_OBJECT_TYPE_KINEMATIC\n"
  "mass: 0.0\n"
  "friction: 0.1\n"
  "restitution: 0.5\n"
  "group: \"player\"\n"
  "mask: \"stone\"\n"
  "mask: \"wall\"\n"
  "embedded_collision_shape {\n"
  "  shapes {\n"
  "    shape_type: TYPE_SPHERE\n"
  "    position {\n"
  "      y: 1.0\n"
  "    }\n"
  "    rotation {\n"
  "    }\n"
  "    index: 0\n"
  "    count: 1\n"
  "  }\n"
  "  data: 5.960093\n"
  "}\n"
  "locked_rotation: true\n"
  ""
}
embedded_components {
  id: "attack"
  type: "factory"
  data: "prototype: \"/player/attacks/Attack.go\"\n"
  ""
}
embedded_components {
  id: "ability"
  type: "factory"
  data: "prototype: \"/player/attacks/Ability.go\"\n"
  ""
}
