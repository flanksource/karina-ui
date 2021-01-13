export type IconsId =
  | "cpu"
  | "memory"
  | "triangle-marker";

export type IconsKey =
  | "Cpu"
  | "Memory"
  | "TriangleMarker";

export enum Icons {
  Cpu = "cpu",
  Memory = "memory",
  TriangleMarker = "triangle-marker",
}

export const ICONS_CODEPOINTS: { [key in Icons]: string } = {
  [Icons.Cpu]: "61697",
  [Icons.Memory]: "61698",
  [Icons.TriangleMarker]: "61699",
};
