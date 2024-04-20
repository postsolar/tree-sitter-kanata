/// <reference types="tree-sitter-cli/dsl" />
// @ts-check

module.exports = grammar({
  name: "kanata",

  extras: $ => [
    $.line_comment,
    $.block_comment,
    /\n/,
    /\r/,
    /\p{Zs}/
  ],

  // TODO
  // regex queries here are very naive with regards to kanata syntax

  rules: {
    kanata: $ => repeat($.list),

    line_comment: _ => seq(';;', /.*/),

    block_comment: _ => seq('#|', /[.\n\r]*/, '|#'),

    list: $ => seq('(', repeat($._item), ')'),

    _item: $ => choice(
      $.unquoted_item,
      $.quoted_item,
      $.list
    ),

    unquoted_item: _ => /[^\s)("]+/,

    quoted_item: _ => seq('"', /[^"]*/, '"')
  }
});
