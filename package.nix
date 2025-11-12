# First, the named parameters.
{
  # you almost always want to depend on lib
  lib,

  # these are Nixpkgs functions that your package uses
  buildGoModule,
  fetchFromGitHub,

  # more dependencies would go here...
}: # this means end of named parameters

# Now, the definition of your package.
# This should be something that produces a derivation, not
# a string or a raw attribute set or anything else.
# buildGoModule is a function that returns a derivation, so
# you want `buildGoModule ...` here, not `{ pet = ...; }` here;
# the latter is an attribute set.

buildGoModule rec {
  pname = "dndbot";
  version = "1.0";

  src = fetchFromGitHub {
    owner = "NilsKrause";
    repo = "dnd";
    tag = "v${version}";
    #hash = "";
  };

  # this hash is updated from the example, which seems to be out of date
  vendorHash = "sha256-6hCgv2/8UIRHw1kCe3nLkxF23zE/7t5RDwEjSzX3pBQ=";

  meta = {
    description = "Simple DnD Discord Bot";
    homepage = "https://github.com/NilsKrause/dnd";
    license = lib.licenses.mit;
    maintainers = with lib.maintainers; [ NilsKrause ];
  };
}
