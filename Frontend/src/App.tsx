import { useEffect, useRef, useState } from 'react'
import vyper from './assets/1.webp'

import './App.css'

function App() {
  // const [count, setCount] = useState(0)
  const [enabled, setEnabled] = useState(false);
  const ref = useRef<HTMLImageElement>(null);


  useEffect(() => {
    if (!enabled) return;

    const id = setInterval(() => {
      const img = ref.current;
      if (img === null) return;
      const r = Math.random() * 180 - 90;
      const distx = Math.random() * 10 + 1;
      const disty = Math.random() * 10 + 1;
      img.style.rotate = `${r}deg`
      img.style.scale = `${distx} ${disty}`
    }, 100);
    return () => clearInterval(id);
  }, [enabled])

  return (
    <>
      <section id="center">
        <div>
          <h1>THE OL' GORGON SPECIAL!</h1>
          <p>
            
          </p>
        </div>
        <button
          type="button"
          className="counter"
          onClick={() => {
            setEnabled(!enabled)
            // setCount((count) => Math.random()*0.3);
            // const img = ref.current;
            // if (img === null) return;
            // img.style.setProperty('--s', count.toString())
            // img.style.transform = `scale(${count})`
            // img.style.rotate = '0deg'
          }}
        >

            <img ref={ref} className="vyper" src={vyper} alt="" style={{ zIndex: '-1'}}/> is ready 

        </button>
      </section>

      <div className="ticks"></div>

      {/* <section id="next-steps">
        <div id="docs">
          <svg className="icon" role="presentation" aria-hidden="true">
            <use href="/icons.svg#documentation-icon"></use>
          </svg>
          <h2>Documentation</h2>
          <p>Your questions, answered</p>
          <ul>
            <li>
              <a href="https://vite.dev/" target="_blank">
                Explore Vite
              </a>
            </li>
            <li>
              <a href="https://react.dev/" target="_blank">
                <img className="button-icon" src={vyper} alt="" />
                Learn more
              </a>
            </li>
          </ul>
        </div>
        
      </section> */}

      <div className="ticks"></div>
      <section id="spacer"></section>
    </>
  )
}

export default App
