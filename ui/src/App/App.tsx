import React from 'react'
import { history } from 'helpers'
import { authenticationService } from 'services'
import { Router, Route, Redirect } from 'react-router-dom'
import { LoginPage, LinkList, ProfilePage } from 'pages'
import { Navigation, PrivateRoute } from 'components'
import { User } from 'models/user'
import { Dashboard } from 'pages/Dashboard'

type Props = {}
type State = {
    currentUser: User | null
}

class App extends React.Component<Props, State> {
    constructor(props: Props) {
        super(props)

        this.state = {
            currentUser: null
        }
    }

    componentDidMount() {
        authenticationService.currentUser.subscribe(x => this.setState({ currentUser: x }))
    }

    logout() {
        authenticationService.logout()
        history.push('/admin/login')
    }

    render() {
        const { currentUser } = this.state

        return (
            <Router history={history}>
                <div className="flex flex-col min-h-screen">
                    {currentUser &&
                        <Navigation user={currentUser} />
                    }
                    <Redirect exact from="/admin" to="/admin/dashboard" />
                    <PrivateRoute exact path="/admin/dashboard" component={Dashboard} />
                    <PrivateRoute exact path="/admin/links" component={LinkList} />
                    <PrivateRoute exact path="/admin/profile" component={ProfilePage} />
                    <Route exact path="/admin/login" component={LoginPage} />
                </div>
            </Router>
        )
    }
}

export { App }