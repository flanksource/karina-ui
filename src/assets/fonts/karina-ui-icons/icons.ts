export type IconsId =
  | "memory"
  | "triangle-marker";

export type IconsKey =
  | "Memory"
  | "TriangleMarker";

export enum Icons {
  Memory = "memory",
  TriangleMarker = "triangle-marker",
}

export const ICONS_CODEPOINTS: { [key in Icons]: string } = {
  [Icons.Memory]: "61697",
  [Icons.TriangleMarker]: "61698",
};
