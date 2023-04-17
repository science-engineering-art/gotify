
export function AudioPlayer(props: { src: string}) {
  return (
    <audio autoPlay={true} src={props.src} controls></audio>
  );
}
