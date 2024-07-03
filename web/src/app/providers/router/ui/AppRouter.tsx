import { Suspense } from "react"
import { Route, Routes } from "react-router-dom"
import { routeConfig } from "shared/config/routeConfig/routeConfig"
import { RequireAuth } from "./RequireAuth"

export const AppRouter= () => {

    return (
        <Suspense fallback={<div>Loading...</div>}>
            <Routes>
                {Object.values(routeConfig)
                    .map((route) => (
                        <Route
                            key={route.path}
                            path={route.path}
                            element={
                                route.authOnly
                                ? 
                                <RequireAuth>
                                    {route.element}
                                </RequireAuth>
                                :
                                route.element
                            }
                        />
                    ))
                }
            </Routes>
        </Suspense>
    )
}